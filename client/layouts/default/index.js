import React from 'react';
import { string, element, oneOfType } from 'prop-types';
import { Container } from './styles';

const DefaultLayout = ({ children }) => <Container>{children}</Container>;

DefaultLayout.propTypes = {
    children: oneOfType([string, element]).isRequired,
};

export default DefaultLayout;
