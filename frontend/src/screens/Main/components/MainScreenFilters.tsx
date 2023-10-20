import {
  Grid,
  Typography,
  Paper,
  TextField,
  Button,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
} from "@mui/material";
import { BotDataFilters } from "../../../services/logs";

interface FilterProps {
  filters: BotDataFilters;
  updateFilters: (filters: Partial<BotDataFilters>) => void;
  onSubmit: () => void;
  players: number[];
}

function MainScreenFilters(props: FilterProps) {
  const { filters, updateFilters, onSubmit, players } = props;

  return (
    <Paper elevation={3} style={{ padding: "16px" }}>
      <Grid container spacing={2}>
        <Grid item xs={12} sm={4}>
          <TextField
            fullWidth
            type="date"
            label="Start date"
            variant="outlined"
            value={filters.startDate}
            onChange={(e) => updateFilters({ startDate: e.target.value})}
          />
        </Grid>
        <Grid item xs={12} sm={4}>
          <TextField
            fullWidth
            type="date"
            label="End date"
            variant="outlined"
            value={filters.endDate}
            onChange={(e) => updateFilters({ endDate: e.target.value})}
          />
        </Grid>
        <Grid item xs={4}>
          <FormControl fullWidth variant="outlined">
            <InputLabel id="player-select-label">Player</InputLabel>
            <Select
              labelId="player-select-label"
              label="Jogador"
              value={filters.player}
              onChange={(e) => updateFilters({ player: Number(e.target.value)})}
            >
              <MenuItem value="">All</MenuItem>
              {players.map(player => (
                <MenuItem value={player}>Player {player}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <Button
            variant="contained"
            color="primary"
            onClick={onSubmit}
          >
            Filter
          </Button>
        </Grid>
      </Grid>
    </Paper>
  );
};

export default MainScreenFilters;
