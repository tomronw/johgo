const searchReducer = (state, action) => {

    switch (action.type) {
        case "FETCH_PRODUCTS":
            return {
                ...state,
                products: action.payload,
                loading: false,
            }
        case "SET_LOADING":
            return {
                ...state,
                loading: true
            }
        case "CLEAR_PRODUCTS":
            return {
                ...state,
                products: []
            }
        default:
            return state;
    }
}

export default searchReducer