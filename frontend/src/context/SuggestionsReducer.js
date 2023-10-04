const suggestionsReducer = (state, action) => {

    switch (action.type) {
        case "FETCH_SUGGESTIONS":
            return {
                ...state,
                suggestions: action.payload,
            }

        default:
            return state;
    }
}

export default suggestionsReducer