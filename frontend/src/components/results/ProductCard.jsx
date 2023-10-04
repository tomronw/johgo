import * as React from 'react';
import '../../css/App.css'

export default function ProductCard({ product }) {
    const isMobile = window.innerWidth <= 768;
    let displayColor;

    if (isMobile) {
        displayColor = '#42f575';
    } else {
        displayColor = '#5691f0';
    }

    return (
        <div className="product-card" style={{ maxWidth: '300px' }}>
            <img
                src={product.image}
                style={{
                    width: 150,
                    height: 150,
                    borderTopLeftRadius: '2px',
                    borderTopRightRadius: '2px',
                    borderRadius: 18,
                    marginBottom: '-5px'
                }}
                alt="title"
            />
            <div
                style={{
                    padding: '8px 16px 16px 16px',
                    borderTop: 'none',
                    borderBottomLeftRadius: '2px',
                    borderBottomRightRadius: '2px'
                }}>
                <h2
                    style={{
                        marginBottom: '-8px',
                        fontSize: 15,
                        maxHeight: '2.4em',
                        overflow: 'clip',
                        textOverflow: 'ellipsis'
                    }}>
                    <a
                        href={product.url}
                        target="_blank"
                        rel="noopener noreferrer"
                        style={{
                            color: displayColor,
                            textDecoration: 'none'
                        }}>
                        {product.title}
                    </a>
                </h2>
                <p
                    style={{
                        marginBottom: '-5px',
                        lineHeight: '1.5',
                        fontSize: 16,
                        fontWeight: 'bold',
                        color: 'white'
                    }}>
                    £{product.price}
                </p>
                <p style={{ marginBottom: '0' }}>
                    <a
                        href={product.siteUrl}
                        target="_blank"
                        rel="noopener noreferrer"
                        style={{
                            color: displayColor,
                            textDecoration: 'none'
                        }}>
                        {product.siteName}
                    </a>
                </p>
            </div>
        </div>

    );
}