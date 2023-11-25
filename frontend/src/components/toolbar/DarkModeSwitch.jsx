import React from 'react';
import { keyframes } from '@emotion/react';
import styled from '@emotion/styled';
import { useTheme } from '../../context/ThemeContext';
import { Tooltip } from '@mui/material';
import DarkModeMoon from '../../images/moon_pixel.png';
import LightModeSun from '../../images/pixel_sun.png';

const slideOut = keyframes`
  from {
    transform: translateX(0);
    opacity: 1;
  }
  to {
    transform: translateX(-100%);
    opacity: 0;
  }
`;

const slideIn = keyframes`
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
`;

const IconWrapper = styled.div`
  display: inline-block;
  cursor: pointer;
  position: relative;
  width: 48px;  // Doubled from 24px
  height: 48px; // Doubled from 24px

  & > svg,
  & > img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    animation-duration: 0.3s;
  }

  & > svg.fade-in,
  & > img.fade-in {
    animation-name: ${slideIn};
  }

  & > svg.fade-out,
  & > img.fade-out {
    animation-name: ${slideOut};
  }
`;

const DarkModeSwitch = () => {
  const { darkMode, setDarkMode } = useTheme();

  const handleClick = () => {
    setDarkMode(!darkMode);
  };

  return (
    <IconWrapper onClick={handleClick}>
      {darkMode ? (
        <Tooltip title="Switch to light mode!">
          <img
            src={LightModeSun}
            alt="Dark Mode Moon"
            className={darkMode ? 'fade-in' : 'fade-out'}
          />
        </Tooltip>
      ) : (
        <Tooltip title="Switch to dark mode!">
          <img
            src={DarkModeMoon}
            alt="Dark Mode Moon"
            className={darkMode ? 'fade-out' : 'fade-in'}
          />
        </Tooltip>
      )}
    </IconWrapper>
  );
};

export default DarkModeSwitch;
