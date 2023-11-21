import React from "react";
import Sidenav from "../../Layouts/Sidenav";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { Box, Button, Checkbox, LinearProgress } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import EditIcon from "@mui/icons-material/Edit";
import { useNavigate } from "react-router-dom";
import useProducts from "../../Services/useProducts";

export default function Barang() {
  const navigate = useNavigate();

  const { data: products, isLoading } = useProducts();

  return (
    <Sidenav
      children={
        <div>
          <h1>Barang</h1>
          <div className="my-3">
            <Button
              variant="contained"
              color="primary"
              onClick={() => {
                navigate("/barang/add");
              }}
            >
              <AddIcon /> Add
            </Button>
            <Button variant="contained" color="warning" className="ms-3">
              <EditIcon /> Edit
            </Button>
          </div>
          <TableContainer component={Paper}>
            <Table
              sx={{ minWidth: 650 }}
              size="small"
              aria-label="table-barang"
            >
              <TableHead>
                <TableRow>
                  <TableCell width={100}>
                    <Checkbox
                      color="primary"
                      // indeterminate={
                      //   numSelected > 0 && numSelected < rowCount
                      // }
                      // checked={rowCount > 0 && numSelected === rowCount}
                      // onChange={onSelectAllClick}
                      // inputProps={{
                      //   "aria-label": "select all desserts",
                      // }}
                    />
                  </TableCell>
                  <TableCell>Code</TableCell>
                  <TableCell>Name</TableCell>
                  <TableCell>Total</TableCell>
                  <TableCell>Description</TableCell>
                  <TableCell>Status</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {products.map((row) => (
                  <TableRow
                    key={row.Id}
                    sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                  >
                    <TableCell component="th" scope="row">
                      <Checkbox
                        color="primary"
                        // indeterminate={
                        //   numSelected > 0 && numSelected < rowCount
                        // }
                        // checked={rowCount > 0 && numSelected === rowCount}
                        // onChange={onSelectAllClick}
                        // inputProps={{
                        //   "aria-label": "select all desserts",
                        // }}
                      />
                    </TableCell>
                    <TableCell>{row.Code}</TableCell>
                    <TableCell>{row.Name}</TableCell>
                    <TableCell>{row.Total}</TableCell>
                    <TableCell>{row.Description}</TableCell>
                    <TableCell>
                      {row.IsActive === 1 ? "Active" : "In Active"}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          {isLoading && (
            <Box sx={{ width: "100%" }}>
              <LinearProgress />
            </Box>
          )}
        </div>
      }
    />
  );
}
