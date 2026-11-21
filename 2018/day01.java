void main() {
    var changes = new BufferedReader(new InputStreamReader(System.in)).lines().mapToInt(Integer::parseInt).toArray();
    IO.print(Arrays.stream(changes).sum());
    int freq = 0;
    var freqs = new HashSet<Integer>();
    freqs.add(freq);
    for (int i = 0; true; i++) {
        freq += changes[i % changes.length];
        if (!freqs.add(freq)) {
            IO.println(" " + freq);
            break;
        }
    }
}