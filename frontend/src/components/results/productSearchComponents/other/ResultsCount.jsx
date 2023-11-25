import React from 'react';

function ResultsCount({ resultsCount }) {

    return (
        <div style={{ color: 'white', fontWeight: 700 }} >
            {resultsCount >= 2 ? <h4>Showing {resultsCount} results</h4> : <h4>0 results</h4>}
        </div>
    );
}

export default ResultsCount;