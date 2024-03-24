package androidx.databinding;

public class DataBinderMapperImpl extends MergedDataBinderMapper {
  DataBinderMapperImpl() {
    addMapper(new org.csystem.android.app.generator.random.DataBinderMapperImpl());
  }
}
