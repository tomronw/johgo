import React from 'react';
import { FormControlLabel, Switch } from '@mui/material';
import "../../css/SearchBox.css";
import Tooltip from '@material-ui/core/Tooltip';
import { useTheme } from '../../context/ThemeContext';

function ExcludeSingles({ checked, handleChange, isMobile }) {
    const { darkMode } = useTheme();
    const fontSizeClass = isMobile ? 'smallFont' : 'largeFont';

    return (
        <Tooltip title="Attempts to filter out all single cards from the search results." placement="left">
            <FormControlLabel
                className={fontSizeClass}
                control={
                    <Switch
                        checked={checked}
                        onChange={handleChange}
                        sx={{
                            '& .MuiSwitch-colorSecondary.Mui-checked': {
                                color: darkMode ? 'yellow' : 'primary',
                            },
                            '& .MuiSwitch-track': {
                                backgroundColor: darkMode ? 'yellow' : 'default',
                            },
                            fontSize: '0.1rem',
                        }}
                    />
                }
                label={<strong style={{ fontSize: isMobile ? '0.7rem' : '0.9rem' }}>Exclude Singles?</strong>}
                labelPlacement="end"
                style={{
                    color: 'white',
                }}
            />
        </Tooltip>
    );
}

export default ExcludeSingles;