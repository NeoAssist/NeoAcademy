import React from 'react';
import { string, element, oneOfType } from 'prop-types';
import { Container } from './styles';

const ErrorLayout = ({ children }) => <Container>{children}</Container>;

ErrorLayout.propTypes = {
    children: oneOfType([string, element]).isRequired,
};

export default ErrorLayout;
