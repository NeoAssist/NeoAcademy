import React from 'react';
import { element } from 'prop-types';
import { Container } from './styles';

const ErrorLayout = ({ children }) => <Container>{children}</Container>;

ErrorLayout.propTypes = {
    children: element.isRequired,
};

export default ErrorLayout;
