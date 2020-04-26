import React from 'react';
import { render } from '@testing-library/react';
import { toBeInTheDocument } from '@testing-library/jest-dom/matchers';
import Dashboard from '../pages/dashboard';

expect.extend({ toBeInTheDocument });

test('renders login', () => {
    const { getByText } = render(<Dashboard />);
    const linkElement = getByText(/Dashboard/);
    expect(linkElement).toBeInTheDocument();
});
