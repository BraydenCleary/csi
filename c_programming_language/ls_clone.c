// Completion exercise:
// Please implement a minimal clone of the ls program. We have chosen this
// exercise as it will require you to use structs, pointers and arrays, as
// well as some C standard library functions with interesting interfaces.
// It will also likely to be substantial enough to merit some degree of code
// organization. Minimally, it should list the contents of a directory including
// some information about each file, such as file size. As a stretch goal, use man
//  ls to identify any interesting flags you may wish to support, and implement them.

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <dirent.h>
#include <sys/stat.h>

#define MAX_FILES 100

// TODO: Expand this struct to collect more info on files
typedef struct {
    long size;
    // TODO: Figure out how to make this a readable string
    long last_modified_at_seconds;
    char *name;
} File;

char *get_cwd(void);
int read_cwd(DIR *cwd, File files[]);
void print_files(File files[], int file_count);

// TODO: Parse args so that you can support ls of any directory, not just cwd
int main() {
    char *current_working_directory = get_cwd();

    DIR *cwd = opendir(current_working_directory);

    if (cwd == NULL) {
        perror("opendir error");
        return 1;
    }

    File files[MAX_FILES];

    int file_count = read_cwd(cwd, files);

    print_files(files, file_count);

    // Do I have to clean up my memory on the heap?
    free(current_working_directory);
    return 0;
}

void print_files(File files[], int file_count) {
    // TODO: Figure out how to make output prettier
    // TODO: support -a flag by filtering out (.) files unless -a passed
    for (int i = 0; i < file_count; i++) {
        printf("Name: %s, Size: %ld bytes, Last Modified At: %lu\n", files[i].name, files[i].size, files[i].last_modified_at_seconds);
    }
}

int read_cwd(DIR *cwd, File files[]) {
    struct dirent *dp;
    int file_count = 0;

    while ((dp = readdir(cwd)) != NULL) {
        struct stat stat_buf;
        int stat_return = stat(dp->d_name, &stat_buf);

        if (stat_return != 0) {
            perror("stat error");
            continue;
        }

        File file = {
            stat_buf.st_size,
            stat_buf.st_mtimespec.tv_sec,
            dp->d_name
        };

        files[file_count] = file;
        file_count++;
    }

    return file_count;
}

char *get_cwd(void) {
    int max_cwd_length = 100;

    // Where should I declare this char buff?
    char *buff = (char*)calloc(max_cwd_length, sizeof(char));

    if (getcwd(buff, max_cwd_length) == NULL) {
        perror("getcwd() error");
        // What's the proper behavior here?
        return NULL;
    }

    return buff;
}
