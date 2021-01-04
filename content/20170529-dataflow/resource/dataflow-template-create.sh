mvn compile exec:java -Dexec.mainClass=org.sinmetal.flow.BigQueryToDatastore \
              -Dexec.args="--project=cpb101demo1 \
              --stagingLocation=gs://cpb101demo1/staging \
              --inputTable=cpb101demo1:samples.table \
              --dataflowJobFile=gs://cpb101demo1-df-template/BigQueryToDatastore201702021500 \
              --runner=TemplatingDataflowPipelineRunner"