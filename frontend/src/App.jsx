import { useEffect, useState } from "react"
import FileInput from "./components/FileInput.jsx"
import Papa from "papaparse"
import Input from "./components/Input.jsx"
import Button from "./components/Button.jsx"
import DropdownCheckList from "./components/DropdownCheckList.jsx"
import DropdownMenu from "./components/DropdownMenu.jsx"
import Status from "./components/Status.jsx"

const initialInputs = {
  host: "jqii6ig8f6.asia-southeast1.gcp.clickhouse.cloud",
  port: "9440",
  database: "default",
  username: "default",
  password: "8ee_Rt6D~CWK7",
  tableName: ""
}

function App() {
  const [columns, setColumns] = useState([])
  const [file, setFile] = useState(null)
  const [selectedColumns, setSelectedColumns] = useState({})
  const [inputs, setInputs] = useState(initialInputs)
  const [source, setSource] = useState("File")
  const [tables, setTables] = useState([])
  const [selectedTable, setSelectedTable] = useState("Select Table")
  const [fileActions, setFileActions] = useState(false)
  const [dbActions, setDbActions] = useState(false)
  const [status, setStatus] = useState("Status: None")
  const [statusType, setStatusType] = useState("")

  useEffect(() => {
    getColumns(file, initializeColumns, setColumns, setSelectedColumns)
    if (file) {
      setStatus("Uploaded File")
      setStatusType("complete")
    }
  }, [file])

  useEffect(() => {
    console.log("Selected columns")
    console.log(selectedColumns)
  }, [selectedColumns])

  useEffect(() => {
    if (source === "ClickHouseDB") {
      setDbActions(true)
      setFileActions(false)
    } else {
      setDbActions(false)
      setFileActions(true)
    }
  }, [source])

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

      <DropdownMenu
        title={"Source: " + source}
        items={["File", "ClickHouseDB"]}
        onClick={(e) => setSource(e.target.id)}
      />

      <VariableDropdown
        source={source}
        selectedColumns={selectedColumns}
        handleColumnSelection={handleColumnSelection}
        tables={tables}
        selectedTable={selectedTable}
        setSelectedTable={setSelectedTable}
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
        disabled={dbActions}
        placeholder={"Table Name"}
        value={inputs.tableName}
        id={"tableName"}
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

      <div>
        <Button
          // disabled={dbActions}
          text={"Connect"}
          onClick={() => connect(inputs, setStatus, setStatusType)}
        ></Button>

        <Button
          disabled={dbActions}
          text={"Upload"}
          onClick={() => upload(columns, selectedColumns, inputs, file, setStatus, setStatusType)}
        ></Button>

        <Button
          disabled={fileActions}
          text={"Fetch Tables"}
          onClick={() => fetchTables(inputs, setTables, setStatus, setStatusType)}
        />

        <Button
          disabled={fileActions}
          text={"Download"}
          onClick={() => fetchRows(inputs, selectedTable, setStatus, setStatusType)}
        ></Button>
      </div>
      <Status
        content={status}
        type={statusType}
      />
    </div>
  )
}

function VariableDropdown({ source, selectedColumns, handleColumnSelection, tables, selectedTable, setSelectedTable }) {
  if (source === 'File') {
    console.log("Selected columns in variable:")
    console.log(selectedColumns)
    return (
      <DropdownCheckList
        items={selectedColumns}
        onChange={handleColumnSelection}
      />
    )
  } else {
    return (
      <DropdownMenu
        title={"Table: " + selectedTable}
        items={tables}
        onClick={(e) => setSelectedTable(e.target.id)}
      />
    )
  }
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

async function fetchTables(inputs, setTables, setStatus, setStatusType) {
  setStatus("Fetching tables")
  setStatusType("progress")

  const data = {
    ConnectionInfo: {
      Host: inputs.host,
      Port: parseInt(inputs.port),
      Database: inputs.database,
      Username: inputs.username,
      Password: inputs.password
    }
  }

  const res = await fetch('http://localhost:8080/get-tables', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  })

  const resJson = await res.json()

  if (res.ok) {
    setStatus("Fetched tables successfuly")
    setStatusType("complete")
  } else {
    setStatus(res.message)
    setStatusType("error")
  }

  setTables(resJson.tables)

  console.log(resJson)
}

async function connect(inputs, setStatus, setStatusType) {
  console.log(inputs)

  setStatus("Connecting...")
  setStatusType("progress")

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

  if (res.ok) {
    setStatus("Connected successfuly")
    setStatusType("complete")
  } else {
    setStatus(resJson.message)
    setStatusType("error")
  }

  console.log(resJson)
}


async function upload(columns, selectedColumnsObj, inputs, file, setStatus, setStatusType) {
  const queue = []
  let processing = false
  let uploadCount = 0

  setStatus("Uploading...")
  setStatusType("progress")

  const selectedColumns = Object.entries(selectedColumnsObj)
    .filter(([_, checked]) => checked)
    .map(([key]) => key)

  const uploadChunk = async function () {
    console.log("Trying")
    if (queue.length == 0) {
      setStatus("Upload successful, uploaded " + uploadCount + " records")
      setStatusType("complete")

      processing = false
      return
    }

    processing = true
    const chunkInfo = queue.shift()

    const chunkNumber = chunkInfo.chunkNumber
    const chunk = chunkInfo.chunk

    console.log(chunkInfo)
    console.log("Processing chunk number: " + chunkNumber)
    const data = {
      ConnectionInfo: {
        Host: inputs.host,
        Port: parseInt(inputs.port),
        Database: inputs.database,
        Username: inputs.username,
        Password: inputs.password
      },
      TableName: inputs.tableName,
      ColumnNames: selectedColumns,
      Rows: chunk
    }

    const res = await fetch('http://localhost:8080/post', {
      method: 'POST',
      body: JSON.stringify(data),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    console.log("Got chunk number: " + chunkNumber)

    const resJson = await res.json()

    if (res.ok) {
      uploadCount += resJson.count
      setStatus("Uploaded " + uploadCount + " records")
    } else {
      setStatus(resJson.message)
      setStatusType("error")
    }

    console.log(resJson)
    uploadChunk()
  }

  let chunkNumber = 0

  Papa.parse(file, {
    header: false,
    // dynamicTyping: true,
    skipEmptyLines: true,
    worker: true,
    chunkSize: 5000000,
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

      const selectedIndices = selectedColumns.map(col => {
        const index = columns.indexOf(col)
        if (index !== -1) {
          return index
        }
      })

      const filteredData = data.map(row => {
        return selectedIndices.map(idx => row[idx])
      })

      console.log(filteredData)

      queue.push({ chunkNumber: chunkNumber, chunk: filteredData })

      if (processing == false) {
        uploadChunk()
      }
    },
    complete: () => {
      if (processing == false) {
        uploadChunk()
      }
      console.log("CSV upload complete.")
    }
  })
}

const fetchRows = async (inputs, selectedTable, setStatus, setStatusType) => {
  let start = 0
  let header = []
  const csvData = []

  let hasMoreRows = true

  setStatus("Fetching rows...")
  setStatusType("progress")

  while (hasMoreRows) {
    try {
      const data = await bringRows(inputs, selectedTable, start, setStatus, setStatusType);

      console.log(data)

      if (start === 0) {
        header = data.columnNames;
      }

      if (!data.rows) {
        hasMoreRows = false;
        setStatus("Preparing download...")
      } else {
        csvData.push(...data.rows)
        setStatus("Ingested " + csvData.length + " rows")
        start = start + data.items
      }
    } catch (e) {
      console.log(e)
      // setStatus(e)
      setStatusType("error")
      return
    }
  }


  console.log(csvData)

  downloadCsv(csvData, header, setStatus, setStatusType)
}

function downloadCsv(csvData, columnHeaders, setStatus, setStatusType) {
  // Stream CSV export in chunks to avoid memory overload
  const csvStream = Papa.unparse({
    fields: columnHeaders,
    data: csvData,
  });

  // Create a Blob with the CSV content
  const blob = new Blob([csvStream], { type: "text/csv;charset=utf-8" });
  const link = document.createElement("a");
  link.href = URL.createObjectURL(blob);
  link.download = "data.csv";
  link.click();

  setStatus("Download complete, ingested " + csvData.length + " records")
  setStatusType("complete")
};

async function bringRows(inputs, selectedTable, start, setStatus, setStatusType) {
  const data = {
    ConnectionInfo: {
      Host: inputs.host,
      Port: parseInt(inputs.port),
      Database: inputs.database,
      Username: inputs.username,
      Password: inputs.password,
    },
    TableName: selectedTable,
    Start: start,  // Offset for pagination
  };

  const res = await fetch("http://localhost:8080/get-rows", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
    },
  });

  const resJson = await res.json();

  if (!res.ok) {
    setStatus(resJson.message)
    setStatusType("Error")
    throw new Error(resJson.message)
  }

  return resJson
}

export default App
