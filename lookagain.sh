#! /bin/bash
find . -type f -name '*.sh' | sort -nr | xargs -I {} basename {} ".sh"