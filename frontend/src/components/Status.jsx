/**
 * Status component renders a styled text based on the `type` prop. 
 * The `content` is displayed with different colors depending on the status type (e.g., "error", "progress", "complete").
 * 
 * @param {Object} props - The properties passed to the Status component.
 * @param {string} [props.content=""] - The text content to display within the span element.
 * @param {string} [props.type=""] - The type of the status. It determines the text color:
 *   - "error" - Red color.
 *   - "progress" - Yellow color.
 *   - "complete" - Green color.
 *   - Default - Gray color.
 * 
 * @returns {JSX.Element} The rendered span element with dynamic text color based on the status type.
 */
function Status({ content = "", type = "" }) {
    let classes

    // Determine the color class based on the status type
    if (type === "error") {
        classes = "text-red-600"
    } else if (type === "progress") {
        classes = "text-yellow-600"
    } else if (type === "complete") {
        classes = "text-green-600"
    } else {
        classes = "text-gray-600"
    }

    return (
        <span className={`${classes} font-poppins font-light`}>{content}</span>
    )
}

export default Status
