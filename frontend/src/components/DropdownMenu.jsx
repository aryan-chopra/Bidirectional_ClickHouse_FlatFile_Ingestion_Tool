import { useState } from "react"

function DropdownMenu({ items = [], onClick, title }) {
    const [open, setOpen] = useState(false)

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

    const toggleDropdown = () => {
        setOpen(prev => !prev)
    }

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
