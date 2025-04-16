import { useState } from "react"

/**
 * DropdownMenu component renders a button that opens a dropdown menu with clickable options.
 * When an option is selected, the `onClick` callback is triggered, and the dropdown is closed.
 * The dropdown items are passed as an array, and the title for the button is also customizable.
 * 
 * @param {Object} props - The properties passed to the DropdownMenu component.
 * @param {Array} props.items - An array of strings representing the menu items to display in the dropdown.
 * @param {function} props.onClick - A callback function that is triggered when a dropdown option is clicked.
 * @param {string} props.title - The text to display on the button that opens the dropdown.
 * 
 * @returns {JSX.Element} The rendered dropdown menu component.
 */
function DropdownMenu({ items = [], onClick, title }) {
    // State to track if the dropdown is open or closed
    const [open, setOpen] = useState(false)

    // Create the list of dropdown options by mapping over the `items` array
    const list = items.map((element, index) => {
        return (
            <button
                key={index}
                id={element}
                onClick={(e) => handleOptionClick(e)}
                className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
            >
                {element}
            </button>
        )
    })

    // Toggle the open/close state of the dropdown
    const toggleDropdown = () => {
        setOpen(prev => !prev)
    }

    // Handle option click: close the dropdown and trigger the onClick callback
    const handleOptionClick = (e) => {
        setOpen(false)
        onClick(e)
    }

    return (
        <div className="relative inline-block text-left">
            <div>
                <button
                    onClick={toggleDropdown}
                    type="button"
                    className="inline-flex justify-center w-full rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none"
                >
                    {title}
                    <svg
                        className="-mr-1 ml-2 h-5 w-5"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M19 9l-7 7-7-7" />
                    </svg>
                </button>
            </div>

            {open && (
                <div className="absolute z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5">
                    <div className="py-1">
                        {list}
                    </div>
                </div>
            )}
        </div>
    )
}

export default DropdownMenu
