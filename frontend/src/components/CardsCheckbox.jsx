import React from 'react';
import { FormControlLabel, Switch } from '@mui/material';
import "../css/SearchBox.css";
import Tooltip from '@material-ui/core/Tooltip';


function CardsCheckbox({ checked, handleChange }) {



    return (
        <Tooltip title="Attempts to filter out all single cards from the search results." placement="left">
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