import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.io.Writer;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class Solution {
    public static void main(String[] args) throws IOException {
        final Scanner scanner = new Scanner(new File("input.txt"));
        final int target = scanner.nextInt();
        final Set<Integer> set = new HashSet<>();

        boolean hasTwoNumbers = false;

        while(!hasTwoNumbers && scanner.hasNext()) {
            final int number = scanner.nextInt();
            if (number < target) {
                final int complement = target - number;
                hasTwoNumbers = set.contains(complement);
                set.add(number);
            }
        }
        scanner.close();

        final Writer writer = new FileWriter("output.txt");
        writer.write(hasTwoNumbers ? "1" : "0");
        writer.close();
    }
}
