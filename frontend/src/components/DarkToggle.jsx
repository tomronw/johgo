import React, { useState } from "react";
import DarkModeToggle from "react-dark-mode-toggle";
import "../css/Logo.css";

function DarkModeSwitch({ changeLightMode }) {
    const [isLightMode, setLightMode] = useState(false);

    const ModeHandler = () => {
        const newMode = !isLightMode;
        setLightMode(newMode);
        changeLightMode(newMode);
    };

    return (
        <div className="dark-mode-switch">
            <DarkModeToggle
                onChange={ModeHandler}
                checked={isLightMode}
                size={60}
            />
        </div>
    );
}

export default DarkModeSwitch;