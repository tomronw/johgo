import React from 'react';
import { Button, Dropdown, Space, ConfigProvider, theme } from 'antd';

function FilterDropdown({ filter, onFilterChange }) {
    const onClick = ({ key }) => {
        switch (key) {
            case '1':
                onFilterChange('low-to-high');
                break;
            case '2':
                onFilterChange('high-to-low');
                break;
            case '3':
                onFilterChange('a-z');
                break;
            case '4':
                onFilterChange('z-a');
                break;
            case '5':
                onFilterChange('default');
                break;
            default:
                break;
        }
    };

    const items = [
        {
            key: '1',
            label: (
                <div>Low to High (Price)</div>
            ),
        },
        {
            key: '2',
            label: (
                <div>
                    High to Low (Price)
                </div>
            ),
        },
        {
            key: '3',
            label: (
                <div>
                    A-Z (Title)
                </div>
            ),
        },
        {
            key: '4',
            label: (
                <div>
                    Z-A (Title)
                </div>
            ),
        },
        {
            key: '5',
            label: (
                <div>
                    Search Relevance
                </div>
            ),
        },
    ];

    return (
        <div style={{ width: '110%' }}>
            <ConfigProvider
                theme={{
                    algorithm: theme.darkAlgorithm,

                }}
            >
                <Space direction="vertical">
                    <Dropdown
                        menu={{
                            items,
                            onClick
                        }}
                        placement="bottom"
                    >
                        <Button>Filter by</Button>
                    </Dropdown>
                </Space>
            </ConfigProvider>
        </div>
    );
}

export default FilterDropdown;
