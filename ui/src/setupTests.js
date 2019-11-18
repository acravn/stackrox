import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
// import react-testing-library extensions once for all tests, as recommended at https://github.com/testing-library/jest-dom#usage
import '@testing-library/jest-dom/extend-expect';

/**
 * Fix for test error "matchMedia not present, legacy browsers require a polyfill"
 * https://github.com/akiran/react-slick/issues/742
 */
if (window.matchMedia) {
    window.matchMedia = window.matchMedia;
} else {
    window.matchMedia = () => ({
        matches: false,
        addListener() {},
        removeListener() {}
    });
}

configure({ adapter: new Adapter() });

class Spy {
    spy = null;

    begin() {
        // jest is magically injected by the jest test runner.
        // eslint-disable-next-line no-undef
        this.spy = jest.spyOn(global.console, 'error');
    }

    assertNotCalled() {
        if (!this.spy) {
            throw new Error('Spy not set!');
        }
        // If you see this line called, that means your test is logging a console error.
        // Look at the console error to see what it is.
        // IF the error you're seeing starts with the following, it means you haven't mocked an API request,
        // and an API request is failing.
        // console.error node_modules/jest-environment-jsdom/node_modules/jsdom/lib/jsdom/virtual-console.js:29
        // Error: Error: connect ECONNREFUSED 127.0.0.1:80
        // To debug this, go to src/services/instance.js and uncomment the commented out code,
        // which will help you figure out which API requests are not being mocked.
        // expect is magically injected by the jest test runner.
        // eslint-disable-next-line no-undef
        expect(this.spy).not.toHaveBeenCalled();
        this.spy = null;
    }
}

const spy = new Spy();

global.beforeEach(() => {
    spy.begin();
});

global.afterEach(() => {
    spy.assertNotCalled();
});
