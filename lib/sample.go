package report

// output of mtr --xml
var sample = `
<?xml version="1.0"?>
<MTR SRC="server0" DST="8.8.8.8" TOS="0x0" PSIZE="64" BITPATTERN="0x00" TESTS="10">
    <HUB COUNT="1" HOST="server1">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>   0.3</Avg>
        <Best>   0.3</Best>
        <Wrst>   0.5</Wrst>
        <StDev>   0.0</StDev>
    </HUB>
    <HUB COUNT="2" HOST="server2">
        <Loss> 40.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    4</Drop>
        <Avg>   0.5</Avg>
        <Best>   0.5</Best>
        <Wrst>   0.6</Wrst>
        <StDev>   0.0</StDev>
    </HUB>
    <HUB COUNT="3" HOST="server3">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>   5.2</Avg>
        <Best>   3.4</Best>
        <Wrst>   7.8</Wrst>
        <StDev>   1.1</StDev>
    </HUB>
    <HUB COUNT="4" HOST="core.azc.airebeam.net">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>   7.3</Avg>
        <Best>   5.3</Best>
        <Wrst>  10.3</Wrst>
        <StDev>   1.6</StDev>
    </HUB>
    <HUB COUNT="5" HOST="209-193-107-37.mammothnetworks.com">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>  10.4</Avg>
        <Best>   8.4</Best>
        <Wrst>  12.9</Wrst>
        <StDev>   1.5</StDev>
    </HUB>
    <HUB COUNT="6" HOST="10.1.31.4">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>  41.6</Avg>
        <Best>  29.9</Best>
        <Wrst> 100.6</Wrst>
        <StDev>  21.3</StDev>
    </HUB>
    <HUB COUNT="7" HOST="google.any2.coresite.com">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>  33.5</Avg>
        <Best>  29.5</Best>
        <Wrst>  39.6</Wrst>
        <StDev>   2.6</StDev>
    </HUB>
    <HUB COUNT="8" HOST="216.239.42.249">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>  33.0</Avg>
        <Best>  29.9</Best>
        <Wrst>  36.4</Wrst>
        <StDev>   2.1</StDev>
    </HUB>
    <HUB COUNT="9" HOST="216.239.42.237">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>  33.1</Avg>
        <Best>  29.2</Best>
        <Wrst>  41.1</Wrst>
        <StDev>   3.7</StDev>
    </HUB>
    <HUB COUNT="10" HOST="google-public-dns-a.google.com">
        <Loss>  0.0%</Loss>
        <Snt>    10</Snt>
        <Drop>    0</Drop>
        <Avg>  31.6</Avg>
        <Best>  28.5</Best>
        <Wrst>  41.3</Wrst>
        <StDev>   3.9</StDev>
    </HUB>
</MTR>
`
