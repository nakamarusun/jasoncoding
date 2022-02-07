for fname in *.svg
do
hash=$(echo $RANDOM | md5sum | head -c 4)
scour -i $fname -o ${fname}o --enable-viewboxing --enable-id-stripping --enable-comment-stripping --shorten-ids --indent=none --shorten-ids --shorten-ids-prefix=${hash}
rm $fname
mv ${fname}o $fname
done
