package org.csystem.android.app.generator.random;

import dagger.MembersInjector;
import dagger.internal.DaggerGenerated;
import dagger.internal.InjectedFieldSignature;
import dagger.internal.QualifierMetadata;
import java.util.concurrent.ExecutorService;
import javax.annotation.processing.Generated;
import javax.inject.Provider;

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
public final class MainActivity_MembersInjector implements MembersInjector<MainActivity> {
  private final Provider<ExecutorService> threadPoolProvider;

  public MainActivity_MembersInjector(Provider<ExecutorService> threadPoolProvider) {
    this.threadPoolProvider = threadPoolProvider;
  }

  public static MembersInjector<MainActivity> create(Provider<ExecutorService> threadPoolProvider) {
    return new MainActivity_MembersInjector(threadPoolProvider);
  }

  @Override
  public void injectMembers(MainActivity instance) {
    injectThreadPool(instance, threadPoolProvider.get());
  }

  @InjectedFieldSignature("org.csystem.android.app.generator.random.MainActivity.threadPool")
  public static void injectThreadPool(MainActivity instance, ExecutorService threadPool) {
    instance.threadPool = threadPool;
  }
}
