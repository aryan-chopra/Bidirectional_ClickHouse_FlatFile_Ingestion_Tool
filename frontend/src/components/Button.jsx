/**
 * Button component renders a customizable button with conditional styling 
 * based on the `disabled` prop. The button includes various styles for 
 * normal, hover, focus, and disabled states. The component supports custom 
 * text and an `onClick` handler.
 * 
 * @param {Object} props - The properties passed to the button component.
 * @param {string} props.text - The text to display inside the button.
 * @param {function} props.onClick - The callback function to execute when the button is clicked.
 * @param {boolean} [props.disabled=false] - A boolean that controls whether the button is disabled. Defaults to `false`.
 * 
 * @returns {JSX.Element} The rendered button element.
 */
function Button({ text, onClick, disabled = false }) {
    let classes

    // Set the button styles based on the disabled state
    if (disabled == false) {
        classes = "rounded-full bg-blue-600 py-2 px-4 border border-transparent text-center text-sm text-white transition-all shadow-md hover:shadow-lg focus:bg-blue-700 focus:shadow-none active:bg-blue-700 hover:bg-blue-700 active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none ml-2"
    } else {
        classes = "rounded-full bg-blue-600 bg-opacity-25 py-2 px-4 border border-transparent text-center text-sm text-white transition-all shadow-md hover:shadow-lg focus:bg-blue-700 focus:shadow-none active:bg-blue-700 hover:bg-blue-700 active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none ml-2"
    }

    return (
        <button
            disabled={disabled}
            className={classes}
            type="button"
            onClick={onClick}
        >
            {text}
        </button>
    )
}

export default Button
