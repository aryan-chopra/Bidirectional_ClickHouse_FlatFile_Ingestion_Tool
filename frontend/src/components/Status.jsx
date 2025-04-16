function Status({ content = "", type = "" }) {
    let classes

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
