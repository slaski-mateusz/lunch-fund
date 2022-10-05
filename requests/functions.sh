
function parseOptions() {
    echo "Parsing options"
    while [ $# -gt 0 ]; do
        if [[ $1 == *"--"* ]]; then
            param="${1/--/}"
            eval $(printf "%q=%q" "$param" "$2")
            echo $1 $2
        fi
        shift
    done
}

function loadBodyTemplate {
    ACTION=`echo $0 | sed s/\.sh//g | sed s,\./,,g`
    echo `cat $ACTION.json | tr '\n' ' ' `
}