import * as React from 'react';


export default function ProductItem({ product }) {
    const isMobile = window.innerWidth <= 768;
    let displayColor;

    if(isMobile){
        displayColor = '#42f575';
    }else{
        displayColor = '#5691f0';
    }

    return (
        <div style={{ maxWidth: '300px' }}>
            <img src={product.Image} style={{ width: 150, height: 150, borderTopLeftRadius: '2px', borderTopRightRadius: '2px', borderRadius: 15 }} alt="title" />
            <div style={{ padding: '16px', borderTop: 'none', borderBottomLeftRadius: '2px', borderBottomRightRadius: '2px' }}>
                <h2 style={{ marginBottom: '2px', fontSize: 15, maxHeight: '2.4em', overflow: 'clip', textOverflow: 'ellipsis' }}><a href={product.Url} target="_blank" rel="noopener noreferrer" style={{ color: displayColor, textDecoration: 'none' }}>{product.Title}</a></h2>
                <p style={{ marginBottom: '1px', lineHeight: '1.5', fontSize: 16, fontWeight: 'bold' }}>£{product.Price}</p>
                <p style={{ marginBottom: '0' }}><a href={product.SiteUrl} target="_blank" rel="noopener noreferrer" style={{ color: displayColor, textDecoration: 'none' }}>{product.SiteName}</a></p>
            </div>
        </div>
    );
}