package org.csystem.android.app.generator.random.di.module.scheduler;

import dagger.internal.DaggerGenerated;
import dagger.internal.Factory;
import dagger.internal.Preconditions;
import dagger.internal.QualifierMetadata;
import dagger.internal.ScopeMetadata;
import java.util.concurrent.ExecutorService;
import javax.annotation.processing.Generated;

@ScopeMetadata("javax.inject.Singleton")
@QualifierMetadata
@DaggerGenerated
@Generated(
    value = "dagger.internal.codegen.ComponentProcessor",
    comments = "https://dagger.dev"
)
@SuppressWarnings({
    "unchecked",
    "rawtypes",
    "KotlinInternal",
    "KotlinInternalInJava"
})
public final class ExecutorServiceModule_CreateSchedulerFactory implements Factory<ExecutorService> {
  @Override
  public ExecutorService get() {
    return createScheduler();
  }

  public static ExecutorServiceModule_CreateSchedulerFactory create() {
    return InstanceHolder.INSTANCE;
  }

  public static ExecutorService createScheduler() {
    return Preconditions.checkNotNullFromProvides(ExecutorServiceModule.INSTANCE.createScheduler());
  }

  private static final class InstanceHolder {
    private static final ExecutorServiceModule_CreateSchedulerFactory INSTANCE = new ExecutorServiceModule_CreateSchedulerFactory();
  }
}
