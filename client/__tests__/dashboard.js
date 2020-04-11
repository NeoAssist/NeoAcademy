import React from 'react';
import { render } from '@testing-library/react';
import Dashboard from '../pages/dashboard';
import { toBeInTheDocument } from '@testing-library/jest-dom/matchers';

expect.extend({ toBeInTheDocument });

test('renders login', () => {
    const { getByText } = render(<Dashboard />);
    const linkElement = getByText(/Dashboard/);
    expect(linkElement).toBeInTheDocument();
});
