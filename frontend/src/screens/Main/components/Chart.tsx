import { Paper, Typography } from "@mui/material";

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
