import React from 'react';
import { FormControlLabel, Switch } from '@mui/material';
import "../css/SearchBox.css";
import Tooltip from '@material-ui/core/Tooltip';


function CardsCheckbox({ checked, handleChange }) {



    return (
        <Tooltip title="Reduces spam by filtering out single cards from search results." placement="left">
            <FormControlLabel
                control={<Switch checked={checked} onChange={handleChange} />}
                label="Exclude Single Cards"
                labelPlacement="end"
                className='exclude-box'
            />
        </Tooltip >
    );
}

export default CardsCheckbox;