import React from 'react';
import { render } from '@testing-library/react';
import { toBeInTheDocument } from '@testing-library/jest-dom/matchers';
import Login from '../pages/login';

expect.extend({ toBeInTheDocument });

test('renders login', () => {
    const { getByText } = render(<Login />);
    const linkElement = getByText(/Login/);
    expect(linkElement).toBeInTheDocument();
});
