#include <stdio.h>
#include <string.h>

int main() {
    FILE *fs = fopen("input.txt", "r");
    int target;
    fscanf(fs, "%d", &target);
    unsigned char bitset[1 + target / 8];
	memset(bitset, 0 , sizeof bitset);

    int isFound = 0;

    for (int number; !isFound && fscanf(fs, "%d", &number) == 1;) {
        if (number <= target) {
            const int compliment = target - number;
            isFound = (bitset[compliment / 8] & (1 << (compliment % 8))) > 0;
            bitset[number / 8] |= 1 << (number % 8);
        }
    }
    fclose(fs);

    fs = fopen("output.txt", "w");
    fprintf(fs,"%d", isFound);
    fclose(fs);

    return 0;
}
