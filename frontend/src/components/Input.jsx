function Input({ id, placeholder, value, onChange, disabled = false }) {
    let classes

    if (disabled == false) {
        classes = "w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
    } else {
        classes = "w-full bg-slate-200 pointer-events-none placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
    }

    return (
        <div className="w-full max-w-sm min-w-[200px]">
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
