export function dijksta(grid, start, target) {

}

function dHeap(G, start='A') {
    var shortest_paths = {};
    var visited = {};
    var heap = [];

    for (var node of Object.keys(G)) {
        shortest_paths[node] = Infinity;
        visited[node] = false;
    }

    shortest_paths[start] = 0;
    visited[start] = true;

    heap.push([0, start]);

    while (heap.length > 0) {
        var [distance, node] = heap.shift();
        visited[node] = true;

        for (var edge of G[node]) {
            var [cost, to_node] = edge;

            if (!visited[to_node] && distance + cost < shortest_paths[to_node]) {
                shortest_paths[to_node] = distance + cost;
                heap.push([shortest_paths[to_node], to_node]);
            }
        }
    }

    return shortest_paths;
}