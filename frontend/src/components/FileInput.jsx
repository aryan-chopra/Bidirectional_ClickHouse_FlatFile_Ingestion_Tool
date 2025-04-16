/**
 * FileInput component renders a file input field that allows users to select a file. 
 * It accepts a file input and triggers the `onChange` callback when a file is selected.
 * 
 * @param {Object} props - The properties passed to the FileInput component.
 * @param {function} props.onChange - A callback function that is triggered when a file is selected.
 * @param {File} [props.file] - The file currently selected in the input (optional).
 * 
 * @returns {JSX.Element} The rendered file input element.
 */
function FileInput({ file, onChange }) {
    return (
        <div
        className="mb-3"
        >
            <label
                htmlFor="file_input"
                className="cursor-pointer sr-only">Choose file</label>
            <input type="file"
                id="file_input"
                typeof=".csv"
                onChange={(e) => onChange(e)}
                className="cursor-pointer block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-900 dark:border-neutral-700 dark:text-neutral-400
                file:bg-gray-50 file:border-0
                file:me-4
                file:py-3 file:px-4
                dark:file:bg-neutral-700 dark:file:text-neutral-400"></input>
        </div>
    )
}

export default FileInput
