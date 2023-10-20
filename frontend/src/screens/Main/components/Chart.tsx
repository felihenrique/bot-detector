import { Paper, Typography } from "@mui/material";
import React, { useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

const data = [
  { date: "2023-10-01", value1: 10, value2: 15, value3: 20 },
  { date: "2023-10-02", value1: 12, value2: 18, value3: 25 },
  { date: "2023-10-03", value1: 15, value2: 22, value3: 30 },
  // Adicione mais dados conforme necessário
];

interface ChartProps {
  players: string[];
  data: {
    date: string;
    [key: string]: any;
  }[];
}

const colors = ["#8884d8", "#82ca9d", "#ffc658"];

function Chart(props: ChartProps) {
  const { players, data } = props;

  return (
    <Paper style={{ marginTop: 30, padding: 15 }}>
      <Typography variant="h6">Bot quantity</Typography>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="date" />
          <YAxis />
          <Tooltip />
          <Legend />
          {players.map((player, index) => (
            <Line
              key={`${player}${index}`}
              name={player}
              type="monotone"
              dataKey={player}
              stroke={colors[index] || colors[0]}
              activeDot={{ r: 5 }}
            />
          ))}
        </LineChart>
      </ResponsiveContainer>
    </Paper>
  );
}

export default Chart;
