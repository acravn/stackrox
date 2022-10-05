package resolvers

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/graphql/resolvers/loaders"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	pkgMetrics "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/scoped"
	"github.com/stackrox/rox/pkg/utils"
)

func init() {
	schema := getBuilder()
	utils.Must(
		// Resolvers for fields in storage.NodeComponent are autogenerated and located in generated.go
		// NOTE: This list is and should remain alphabetically ordered
		schema.AddExtraResolvers("NodeComponent", []string{
			"fixedIn: String!",
			"lastScanned: Time",
			"location(query: String): String!",
			"nodes(query: String, scopeQuery: String, pagination: Pagination): [Node!]!",
			"nodeCount(query: String, scopeQuery: String): Int!",
			"nodeVulnerabilities(query: String, scopeQuery: String, pagination: Pagination): [NodeVulnerability]!",
			"nodeVulnerabilityCount(query: String, scopeQuery: String): Int!",
			"nodeVulnerabilityCounter(query: String): VulnerabilityCounter!",
			"plottedNodeVulnerabilities(query: String): PlottedNodeVulnerabilities!",
			"source: String!",
			"topNodeVulnerability: NodeVulnerability",
			"unusedVarSink(query: String): Int",
		}),

		schema.AddQuery("nodeComponent(id: ID): NodeComponent"),
		schema.AddQuery("nodeComponents(query: String, scopeQuery: String, pagination: Pagination): [NodeComponent!]!"),
		schema.AddQuery("nodeComponentCount(query: String): Int!"),
	)
}

// NodeComponentResolver represents a generic resolver of node component fields.
// NOTE: This list is and should remain alphabetically ordered
type NodeComponentResolver interface {
	FixedIn(ctx context.Context) string
	Id(ctx context.Context) graphql.ID
	LastScanned(ctx context.Context) (*graphql.Time, error)
	Location(ctx context.Context, args RawQuery) (string, error)
	Name(ctx context.Context) string
	Nodes(ctx context.Context, args PaginatedQuery) ([]*nodeResolver, error)
	NodeCount(ctx context.Context, args RawQuery) (int32, error)
	NodeVulnerabilities(ctx context.Context, args PaginatedQuery) ([]NodeVulnerabilityResolver, error)
	NodeVulnerabilityCount(ctx context.Context, args RawQuery) (int32, error)
	NodeVulnerabilityCounter(ctx context.Context, args RawQuery) (*VulnerabilityCounterResolver, error)
	OperatingSystem(ctx context.Context) string
	PlottedNodeVulnerabilities(ctx context.Context, args RawQuery) (*PlottedNodeVulnerabilitiesResolver, error)
	Priority(ctx context.Context) int32
	RiskScore(ctx context.Context) float64
	Source(ctx context.Context) string
	TopNodeVulnerability(ctx context.Context) (NodeVulnerabilityResolver, error)
	UnusedVarSink(ctx context.Context, args RawQuery) *int32
	Version(ctx context.Context) string
}

// NodeComponent returns a node component based on an input id (name:version)
func (resolver *Resolver) NodeComponent(ctx context.Context, args IDQuery) (NodeComponentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "NodeComponent")
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		return resolver.nodeComponentV2(ctx, args)
	}

	if err := readNodes(ctx); err != nil {
		return nil, err
	}
	loader, err := loaders.GetNodeComponentLoader(ctx)
	if err != nil {
		return nil, err
	}

	ret, err := loader.FromID(ctx, string(*args.ID))
	res, err := resolver.wrapNodeComponent(ret, true, err)
	if err != nil {
		return nil, err
	}
	res.ctx = ctx
	return res, nil
}

// NodeComponents returns node components that match the input query.
func (resolver *Resolver) NodeComponents(ctx context.Context, q PaginatedQuery) ([]NodeComponentResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "NodeComponents")
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		query := queryWithNodeIDRegexFilter(q.String())

		return resolver.nodeComponentsV2(ctx, PaginatedQuery{Query: &query, Pagination: q.Pagination})
	}

	if err := readNodes(ctx); err != nil {
		return nil, err
	}
	query, err := q.AsV1QueryOrEmpty()
	if err != nil {
		return nil, err
	}
	loader, err := loaders.GetNodeComponentLoader(ctx)
	if err != nil {
		return nil, err
	}

	componentResolvers, err := resolver.wrapNodeComponents(loader.FromQuery(ctx, query))
	if err != nil {
		return nil, err
	}

	ret := make([]NodeComponentResolver, 0, len(componentResolvers))
	for _, res := range componentResolvers {
		res.ctx = ctx
		ret = append(ret, res)
	}
	return ret, nil
}

// NodeComponentCount returns count of node components that match the input query
func (resolver *Resolver) NodeComponentCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.Root, "NodeComponentCount")
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		query := queryWithNodeIDRegexFilter(args.String())

		return resolver.componentCountV2(ctx, RawQuery{Query: &query})
	}

	if err := readNodes(ctx); err != nil {
		return 0, err
	}
	query, err := args.AsV1QueryOrEmpty()
	if err != nil {
		return 0, err
	}
	loader, err := loaders.GetNodeComponentLoader(ctx)
	if err != nil {
		return 0, err
	}

	return loader.CountFromQuery(ctx, query)
}

/*
Utility Functions
*/

func queryWithNodeIDRegexFilter(q string) string {
	return search.AddRawQueriesAsConjunction(q,
		search.NewQueryBuilder().AddRegexes(search.NodeID, ".+").Query())
}

func (resolver *nodeComponentResolver) withNodeComponentScope(ctx context.Context) context.Context {
	if env.PostgresDatastoreEnabled.BooleanSetting() {
		return scoped.Context(ctx, scoped.Scope{
			Level: v1.SearchCategory_NODE_COMPONENTS,
			ID:    resolver.data.GetId(),
		})
	}
	return scoped.Context(ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGE_COMPONENTS,
		ID:    resolver.data.GetId(),
	})
}

func (resolver *nodeComponentResolver) nodeComponentQuery() *v1.Query {
	return search.NewQueryBuilder().AddExactMatches(search.ComponentID, resolver.data.GetId()).ProtoQuery()
}

func (resolver *nodeComponentResolver) nodeComponentRawQuery() string {
	return search.NewQueryBuilder().AddExactMatches(search.ComponentID, resolver.data.GetId()).Query()
}

/*
Sub Resolver Functions
*/

// FixedIn returns the node component version that fixes all the fixable vulnerabilities in this component.
func (resolver *nodeComponentResolver) FixedIn(ctx context.Context) string {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "FixedIn")
	return ""
}

// LastScanned is the last time the node component was scanned in a node.
func (resolver *nodeComponentResolver) LastScanned(_ context.Context) (*graphql.Time, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "LastScanned")
	nodeLoader, err := loaders.GetNodeLoader(resolver.ctx)
	if err != nil {
		return nil, err
	}

	componentQuery := resolver.nodeComponentQuery()
	componentQuery.Pagination = &v1.QueryPagination{
		Limit:  1,
		Offset: 0,
		SortOptions: []*v1.QuerySortOption{
			{
				Field:    search.NodeScanTime.String(),
				Reversed: true,
			},
		},
	}

	nodes, err := nodeLoader.FromQuery(resolver.ctx, componentQuery)
	if err != nil || len(nodes) == 0 {
		return nil, err
	} else if len(nodes) > 1 {
		return nil, errors.New("multiple nodes matched for last scanned component query")
	}

	return timestamp(nodes[0].GetScan().GetScanTime())
}

// Location of the node component.
func (resolver *nodeComponentResolver) Location(_ context.Context, _ RawQuery) (string, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "Location")
	return "Not Available", nil
}

// Nodes that contain the node component.
func (resolver *nodeComponentResolver) Nodes(ctx context.Context, args PaginatedQuery) ([]*nodeResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "Nodes")
	return resolver.root.Nodes(resolver.withNodeComponentScope(ctx), args)
}

// NodeCount is the number of nodes that contain the node component
func (resolver *nodeComponentResolver) NodeCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "NodeCount")
	return resolver.root.NodeCount(resolver.withNodeComponentScope(ctx), args)
}

// NodeVulnerabilities contained in the node component
func (resolver *nodeComponentResolver) NodeVulnerabilities(ctx context.Context, args PaginatedQuery) ([]NodeVulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "NodeVulnerabilities")
	return resolver.root.NodeVulnerabilities(resolver.withNodeComponentScope(ctx), args)
}

// NodeVulnerabilityCount resolves the number of node vulnerabilities contained in the node component
func (resolver *nodeComponentResolver) NodeVulnerabilityCount(ctx context.Context, args RawQuery) (int32, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "NodeVulnerabilityCount")
	return resolver.root.NodeVulnerabilityCount(resolver.withNodeComponentScope(ctx), args)
}

// NodeVulnerabilityCounter resolves the number of different types of node vulnerabilities contained in a node component
func (resolver *nodeComponentResolver) NodeVulnerabilityCounter(ctx context.Context, args RawQuery) (*VulnerabilityCounterResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "NodeVulnerabilityCounter")
	return resolver.root.NodeVulnerabilityCounter(resolver.withNodeComponentScope(ctx), args)
}

// PlottedNodeVulnerabilities returns the data required by top risky component scatter-plot on vuln mgmt dashboard
func (resolver *nodeComponentResolver) PlottedNodeVulnerabilities(ctx context.Context, args RawQuery) (*PlottedNodeVulnerabilitiesResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "PlottedNodeVulnerabilities")
	return resolver.root.PlottedNodeVulnerabilities(resolver.withNodeComponentScope(ctx), args)
}

// Source returns the source type of the node component
func (resolver *nodeComponentResolver) Source(_ context.Context) string {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "Source")
	return storage.SourceType_INFRASTRUCTURE.String()
}

// TopNodeVulnerability returns the first node component vulnerability with the top CVSS score
func (resolver *nodeComponentResolver) TopNodeVulnerability(ctx context.Context) (NodeVulnerabilityResolver, error) {
	defer metrics.SetGraphQLOperationDurationTime(time.Now(), pkgMetrics.NodeComponents, "TopNodeVulnerability")
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		query := resolver.nodeComponentQuery()
		query.Pagination = &v1.QueryPagination{
			SortOptions: []*v1.QuerySortOption{
				{
					Field:    search.CVSS.String(),
					Reversed: true,
				},
				{
					Field:    search.CVE.String(),
					Reversed: true,
				},
			},
			Limit:  1,
			Offset: 0,
		}

		vulnLoader, err := loaders.GetCVELoader(ctx)
		if err != nil {
			return nil, err
		}
		vulns, err := vulnLoader.FromQuery(ctx, query)
		if err != nil || len(vulns) == 0 {
			return nil, err
		} else if len(vulns) > 1 {
			return nil, errors.New("multiple vulnerabilities matched for top node component vulnerability")
		}

		res, err := resolver.root.wrapCVE(vulns[0], true, nil)
		if err != nil {
			return nil, err
		}
		res.ctx = ctx
		return res, nil
	}

	return resolver.root.TopNodeVulnerability(resolver.withNodeComponentScope(ctx), RawQuery{})
}

// UnusedVarSink represents a query sink
func (resolver *nodeComponentResolver) UnusedVarSink(_ context.Context, _ RawQuery) *int32 {
	return nil
}
