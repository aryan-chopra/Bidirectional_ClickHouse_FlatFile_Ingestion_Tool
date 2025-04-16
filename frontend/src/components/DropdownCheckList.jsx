import { useState } from "react"

function DropdownCheckList({ items = {}, onChange }) {
    console.log("Items:")
    console.log(items)
    const handleToggle = (key, currentValue) => {
        const fakeEvent = {
            target: {
                id: key,
                checked: !currentValue,
            },
        }
        onChange(fakeEvent)
    }

    const list = Object.entries(items).map(([key, value]) => (
        <div
            key={key}
            onClick={() => handleToggle(key, value)}
            className="flex items-center gap-x-2 px-3 py-1 rounded-md hover:bg-blue-50 cursor-pointer transition duration-150 ease-in-out"
        >
            <input
                id={key}
                type="checkbox"
                checked={value}
                onChange={(e) => {
                    e.stopPropagation() // prevent double-toggle
                    onChange(e)
                }}
                className="cursor-pointer text-blue-600 h-4 w-4"
            />
            <span className="text-sm text-gray-800 cursor-pointer break-words">
                {key}
            </span>
        </div>
    ))

    console.log("List")
    console.log(list)
    const [isOpen, setIsOpen] = useState(false);

    return (
        <>
            <div className="relative inline-block text-left">
                <button
                    type="button"
                    onClick={() => setIsOpen(!isOpen)}
                    className="w-full rounded-lg border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                    Select Columns
                    <svg
                        className="w-5 h-5 ml-2 inline-block float-right"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                    </svg>
                </button>

                {isOpen && (
                    <div className="absolute z-10 mt-2 w-full origin-top-right rounded-lg bg-white shadow-lg ring-1 ring-black ring-opacity-5">
                        <div className="py-1 px-3 flex flex-col gap-2">
                            {list}
                        </div>
                    </div>
                )}
            </div>
        </>
    )
}

export default DropdownCheckList
