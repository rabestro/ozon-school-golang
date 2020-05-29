import java.util.Set;
import java.util.HashSet;
import java.util.Scanner;

public class Solution {
	public static void main(String[] args) {  
	    final Set<Long> set = new HashSet<>();
	    final Scanner scanner = new Scanner(System.in);
		
		while(scanner.hasNext()) {
		    final Long n = scanner.nextLong();
		    if (set.contains(n)) {
		        set.remove(n);
		    } else {
		        set.add(n);
		    }
		}
		System.out.println(set.iterator().next());
	}
}
