for fname in *.svg
do
# mv $fname $(head -1 -q $fname|awk '{print $1}').inline.svg
TMPA=(${fname//-/ })
mv $fname ${TMPA[0]}.inline.svg
done
