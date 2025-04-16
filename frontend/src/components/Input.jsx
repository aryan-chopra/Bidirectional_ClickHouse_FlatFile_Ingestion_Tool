/**
 * Input component renders a customizable text input field.
 * It supports `disabled` state and triggers the `onChange` callback when the value is modified.
 * It can accept placeholder text, custom values, and manage input styles dynamically based on the `disabled` state.
 * 
 * @param {Object} props - The properties passed to the Input component.
 * @param {string} props.id - The unique identifier for the input field (used for `id` and `htmlFor` attributes).
 * @param {string} props.placeholder - The placeholder text displayed in the input field.
 * @param {string} props.value - The current value of the input field.
 * @param {function} props.onChange - A callback function that is triggered when the value of the input changes.
 * @param {boolean} [props.disabled=false] - Whether the input field is disabled. Default is `false`.
 * 
 * @returns {JSX.Element} The rendered input field component.
 */
function Input({ id, placeholder, value, onChange, disabled = false }) {
    let classes

    // Conditional styling for disabled vs enabled input field
    if (disabled == false) {
        classes = "w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
    } else {
        classes = "w-full bg-slate-200 pointer-events-none placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
    }

    return (
        <div className="w-full min-w-[200px]">
            <input
                disabled={disabled}
                id={id}
                className={classes}
                placeholder={placeholder}
                onChange={(e) => onChange(e)}
                value={value}
            ></input>
        </div>
    )
}

export default Input
