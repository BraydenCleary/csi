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
#include <errno.h>
#include <sys/stat.h>

#define MAX_FILES 400
#define MAX_FILENAME_LENGTH 300

// TODO: Expand this struct to collect more info on files
typedef struct {
    long size;
    // TODO: Figure out how to make this a readable string
    long last_modified_at_seconds;
    char *name;
} File;

char *get_cwd(void);
int read_dir(char *directory_name, DIR *dir, File files[]);
void print_files(File files[], int file_count);
int read_file(char *filename, File files[], int file_count);

// TODO: Parse args so that you can support ls of any directory, not just cwd
int main(int argc, char *argv[]) {
    char *entity_to_open;

    if (argc == 1) {
        entity_to_open = get_cwd();
    } else {
        // If file is passed, print information about file
        // If directory is passed, print information about files in directory
        entity_to_open = argv[1];
    }

    printf("entity to open: %s\n", entity_to_open);

    DIR *dir = opendir(entity_to_open);

    File files[MAX_FILES];
    int file_count = 0;

    if (dir == NULL) {
        // Let's assume that we're working with a file
        file_count = read_file(entity_to_open, files, file_count);
    } else {
        file_count = read_dir(entity_to_open, dir, files);
    }

    print_files(files, file_count);

    // Do I have to clean up my memory on the heap?
    if (argc == 1) {
        free(entity_to_open);
    }
    return 0;
}

int read_file(char *filename, File files[], int file_count) {
    struct stat stat_buf;
    int stat_return = stat(filename, &stat_buf);

    printf("filename: %s\n", filename);

    if (stat_return != 0) {
        fprintf(stderr, "stat error for %s: %s\n", filename, strerror(errno));
    }

    File file = {
        stat_buf.st_size,
        stat_buf.st_mtimespec.tv_sec,
        filename
    };

    files[file_count] = file;

    return file_count + 1;
}

void print_files(File files[], int file_count) {
    // TODO: Figure out how to make output prettier
    // TODO: support -a flag by filtering out (.) files unless -a passed
    for (int i = 0; i < file_count; i++) {
        printf("Name: %s, Size: %ld bytes, Last Modified At: %lu\n", files[i].name, files[i].size, files[i].last_modified_at_seconds);
    }
}

int read_dir(char *directory_name, DIR *dir, File files[]) {
    struct dirent *dp;
    int file_count = 0;

    while ((dp = readdir(dir)) != NULL && file_count < MAX_FILES) {
        char *full_path = (char *)calloc(MAX_FILENAME_LENGTH, sizeof(char));
        strcat(full_path, directory_name);
        strcat(full_path, "/");
        strcat(full_path, dp->d_name);
        file_count = read_file(full_path, files, file_count);
    }

    // Check if more files to read and let user know that we stopped reading files
    // will have to allocate more memory for files array

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
