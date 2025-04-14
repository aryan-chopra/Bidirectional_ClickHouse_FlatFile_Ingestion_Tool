import { useEffect, useState } from "react"
import FileInput from "./components/FileInput.jsx"
import DropdownList from "./components/DropdownList.jsx"
import Papa from "papaparse"

function App() {
  const [columns, setColumns] = useState([])
  const [file, setFile] = useState(null)

  useEffect(() => {
    getColumns(file, setColumns)
  }, [file])

  const handleFileChange = (event) => {
    const file = event.target.files[0]
    setFile(file)
  }

  const getColumns = (file, onComplete) => {
    if (!file) {
      return
    }

    console.log("Parsing")

    Papa.parse(file,
      {
        preview: 1,
        header: true,
        complete: (results) => {
          console.log(results.meta.fields)
          onComplete(results.meta.fields)
        }
      }
    )
  }

  return (
    <div>
      <FileInput
        file={file}
        onChange={handleFileChange}
      />

      <DropdownList
        items={columns}
      />
    </div>
  )
}

export default App
