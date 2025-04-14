import { useEffect, useState } from "react"
import FileInput from "./components/FileInput.jsx"
import DropdownList from "./components/DropdownList.jsx"
import Papa from "papaparse"

function App() {
  const [columns, setColumns] = useState([])
  const [file, setFile] = useState(null)
  const [selectedColumns, setSelectedColumns] = useState({})

  useEffect(() => {
    getColumns(file, initializeColumns)
  }, [file])

  useEffect(() => {
    console.log("Selected columns")
    console.log(selectedColumns)
  },[selectedColumns])

  const initializeColumns = (columns) => {
    setColumns(columns)
    
    let tempColumnObjects = {}

    for (const column of columns) {
      tempColumnObjects[column] = true
    }

    setSelectedColumns(tempColumnObjects)
  }

  const handleFileChange = (event) => {
    const file = event.target.files[0]
    setFile(file)
  }

  const handleColumnSelection = (event) => {
    const id = event.target.id
    const checked = event.target.checked

    console.log("Marking: " + id + " as " + checked)

    setSelectedColumns(selectedColumns => {
      return {...selectedColumns, [`${id}`]: checked}
    })
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
        onChange={handleColumnSelection}
      />
    </div>
  )
}

export default App
