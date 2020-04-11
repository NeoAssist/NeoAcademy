import React from 'react';
import ErrorLayout from '../layouts/error';

const Error = ({ statusCode }) => {
    return (
        <ErrorLayout>
            <div>Deu ruim com esse código => {statusCode}</div>
        </ErrorLayout>
    );
};

Error.getInitialProps = ({ res, err }) => {
    const statusCode = res ? res.statusCode : err ? err.statusCode : 404;
    return { statusCode };
};

export default Error;
