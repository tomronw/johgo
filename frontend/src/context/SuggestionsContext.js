import { createContext, useReducer } from 'react';
import suggestionsReducer from './SuggestionsReducer';

const SuggestionContext = createContext();

export const SuggestionProvider = ({ children }) => {

    const initialState = {
        suggestions: ['', '', '', '', '']
    }

    const [state, dispatch] = useReducer(suggestionsReducer, initialState)

    const fetchSuggestions = async () => {

        try {
            const response = await fetch(process.env.REACT_APP_SUGGESTIONS);

            if (response.ok) {
                const data = await response.json();
                dispatch({ type: "FETCH_SUGGESTIONS", payload: data })
            } else {
                dispatch({ type: "FETCH_SUGGESTIONS", payload: ['', '', '', '', ''] })
            }
        } catch (error) {
            dispatch({ type: "FETCH_SUGGESTIONS", payload: ['', '', '', '', ''] })
        }
    }

    return (
        <SuggestionContext.Provider value={{ suggestions: state.suggestions, fetchSuggestions }}>
            {children}
        </SuggestionContext.Provider>
    );
};

export default SuggestionContext;
