import React from 'react';
import renderer from 'react-test-renderer';
import Login from '../pages/login';

it('renders homepage unchanged', () => {
    const tree = renderer.create(<Login />).toJSON();
    expect(tree).toMatchSnapshot();
});
