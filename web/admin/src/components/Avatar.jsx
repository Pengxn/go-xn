import gravatar from "@xn-02f/gravatar"

export default ({ email }) => {
    return (
        <img src={ gravatar(email) } />
    )
}
