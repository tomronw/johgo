import React from 'react';
import { AppBar, Toolbar, Typography, IconButton } from '@material-ui/core';
import { Link } from 'react-router-dom';
import { createTheme, ThemeProvider, makeStyles } from '@material-ui/core/styles';
import "typeface-cormorant";
import '../css/Index.css'
import logo from '../images/johgoLogo.png'
import '../css/Logo.css'
import DarkModeSwitch from './DarkToggle';

const theme = createTheme({
    typography: {
        fontFamily: ['Segoe UI', 'Roboto', 'Oxygen',
            'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', 'sans-serif'].join(','),
        fontSize: 16,
        allVariants: {
            color: "white"
        },
        fontWeight: 'bold',
    },
});

const isMobile = window.innerWidth <= 768;

const useStyles = makeStyles((theme) => ({
    title: {
        flexGrow: 1,
        fontWeight: 'bold',
    },
    link: {
        color: theme.palette.common.white,
        textDecoration: 'none',
        marginRight: isMobile ? theme.spacing(0) : theme.spacing(5),
        fontWeight: 'bold',
        marginTop: isMobile ? theme.spacing(-9) : theme.spacing(-5),
    },

}));

function MainAppBar({ setLightMode }) {

    const classes = useStyles();

    const changeLightMode = (newMode) => {
        setLightMode(newMode);
    }


    return (
        <ThemeProvider theme={theme} className="body">
            <AppBar position="static" elevation={0} style={{ background: 'transparent', boxShadow: 'none', fontWeight: 'bold' }} className="body">
                <Toolbar>
                    <div>
                        {!isMobile ? <DarkModeSwitch changeLightMode={changeLightMode} /> : null}
                    </div>
                    <div style={{
                        position: 'relative',
                    }} className="logo-container">
                        <a href="/" style={{ textDecoration: 'none' }}>
                            <img src={logo} alt="Logo" />
                        </a>
                    </div>
                    <Typography variant="h6" className={classes.title}>
                    </Typography>
                    <Link to="/about" className={classes.link}>
                        <IconButton>
                            <Typography className={classes.link}>
                                About
                            </Typography>
                        </IconButton>
                    </Link>
                </Toolbar>
            </AppBar>
        </ThemeProvider>
    );
}

export default MainAppBar;