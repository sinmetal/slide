public interface BigQueryToDatastoreOptions extends PipelineOptions {

    @Description("Path of the bigquery table to read from")
    @Default.String("cpb101demo1:samples.table")
    ValueProvider<String> getInputTable();

    void setInputTable(ValueProvider<String> value);
}
