import React from 'react';
import { AppBar, Toolbar, Typography, Box } from '@mui/material';
import DarkModeSwitch from './DarkModeSwitch';
import JohGoLogo from '../../images/johgoLogo.png'
import '../../css/Logo.css'

function Mainbar() {
    return (
        <AppBar position="static" color="transparent" elevation={0}>
            <Toolbar style={{ position: 'relative' }}>
                <Box style={{ position: 'absolute', left: '16px', top: '15%', transform: 'translateY(0%)' }}>
                    <DarkModeSwitch />
                </Box>
                <Typography variant="h6" style={{ marginLeft: 'auto', marginRight: 'auto', marginTop: '2%' }}>
                    <a href="/" style={{ textDecoration: 'none' }}>
                        <img
                            src={JohGoLogo}
                            alt="JohGo Logo"
                            style={{
                                width: '40vw',
                                maxWidth: '300px',
                                minWidth: '100px',
                                padding: '10px 0',
                                animation: 'float 2s infinite',
                                border: 'none',
                                outline: 'none'
                            }}
                        />
                    </a>
                </Typography>


            </Toolbar>
        </AppBar>
    );
}

export default Mainbar;



