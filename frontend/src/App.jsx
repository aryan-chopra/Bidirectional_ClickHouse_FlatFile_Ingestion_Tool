import { useEffect, useState } from "react"
import FileInput from "./components/FileInput.jsx"
import DropdownList from "./components/DropdownList.jsx"
import Papa from "papaparse"
import Input from "./components/Input.jsx"
import Button from "./components/Button.jsx"

const initialInputs = {
  host: "",
  port: "",
  database: "",
  username: "",
  password: ""
}

function App() {
  const [columns, setColumns] = useState([])
  const [file, setFile] = useState(null)
  const [selectedColumns, setSelectedColumns] = useState({})
  const [inputs, setInputs] = useState(initialInputs)

  useEffect(() => {
    getColumns(file, initializeColumns, setColumns, setSelectedColumns)
  }, [file])

  useEffect(() => {
    console.log("Selected columns")
    console.log(selectedColumns)
  }, [selectedColumns])


  const handleFileChange = (event) => {
    const file = event.target.files[0]
    setFile(file)
  }

  const handleColumnSelection = (event) => {
    const id = event.target.id
    const checked = event.target.checked

    console.log("Marking: " + id + " as " + checked)

    setSelectedColumns(selectedColumns => {
      return { ...selectedColumns, [`${id}`]: checked }
    })
  }

  const handleInputChange = (event) => {
    const id = event.target.id
    const value = event.target.value

    setInputs(inputs => {
      return { ...inputs, [id]: value }
    })
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

      <Input
        placeholder={"Hostname"}
        value={inputs.host}
        id={"host"}
        onChange={handleInputChange}
      >
      </Input>

      <Input
        placeholder={"Port"}
        value={inputs.port}
        id={"port"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Database"}
        value={inputs.database}
        id={"database"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Username"}
        value={inputs.username}
        id={"username"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Password"}
        value={inputs.password}
        id={"password"}
        onChange={handleInputChange}
      ></Input>

      <Button
        text={"Connect"}
        onClick={() => connect(inputs)}
      ></Button>

      <Button
        text={"Upload"}
        onClick={() => upload(file)}
      ></Button>
    </div>
  )
}

function getColumns(file, onComplete, setColumns, setSelectedColumns) {
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
        onComplete(results.meta.fields, setColumns, setSelectedColumns)
      }
    }
  )
}

function initializeColumns(columns, setColumns, setSelectedColumns) {
  setColumns(columns)

  let tempColumnObjects = {}

  for (const column of columns) {
    tempColumnObjects[column] = true
  }

  setSelectedColumns(tempColumnObjects)
}

async function connect(inputs) {
  console.log(inputs)

  const data = {
    Host: inputs.host,
    Port: parseInt(inputs.port),
    Database: inputs.database,
    Username: inputs.username,
    Password: inputs.password
  }

  console.log(data)

  const res = await fetch('http://localhost:8080/connect', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  })

  const resJson = await res.json()

  console.log(resJson)
}

async function upload(file) {
  let chunkNumber = 0

  Papa.parse(file, {
    header: false,
    skipEmptyLines: true,
    worker: true,
    chunkSize: 300, // Papa handles chunking!
    chunk: async (results, parser) => {
      let data

      if (chunkNumber == 0) {
        data = results.data.slice(1)
      } else {
        data = results.data
      }
      chunkNumber++
      console.log("Chunk:")
      console.log(data)
    },
    complete: () => {
      console.log("CSV upload complete.")
    }
  })
}

export default App
