class Solution {
public:
    int countPaths(vector<vector<int>>& grid) {
        int unique_path = 0;
        vector<vector<int>> visit (grid.size(), vector<int>(grid[0].size(), 0));
        unique_path = dfs(grid, 0, 0, visit);
        return unique_path;
    }

    int dfs(vector<vector<int>>& grid, int r, int c, vector<vector<int>>& visit){
        int ROWS = grid.size(), COLS = grid[0].size();

        // Check the base case.
        if (min(r, c) < 0 || r == ROWS || c == COLS || visit[r][c] || grid[r][c] == 1){
            return 0;
        }
        // Find a path 
        if (r == ROWS - 1 && c == COLS - 1){
            return 1;
        }

        int count = 0;
        // Update visit
        visit[r][c] = 1;

        // Recursive call
        count += dfs(grid, r - 1, c, visit);
        count += dfs(grid, r + 1, c, visit);
        count += dfs(grid, r, c - 1, visit);
        count += dfs(grid, r, c + 1, visit);

        // Reset visit
        visit[r][c] = 0;

        return count;
    }
};
