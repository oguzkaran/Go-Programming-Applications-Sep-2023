package org.csystem.android.app.generator.random.di.module.scheduler;

import dagger.Module;
import dagger.Provides;
import dagger.hilt.InstallIn;
import dagger.hilt.components.SingletonComponent;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import javax.inject.Singleton;

@dagger.Module
@kotlin.Metadata(mv = {1, 9, 0}, k = 1, xi = 48, d1 = {"\u0000\u0012\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\b\u00c7\u0002\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002J\b\u0010\u0003\u001a\u00020\u0004H\u0007\u00a8\u0006\u0005"}, d2 = {"Lorg/csystem/android/app/generator/random/di/module/scheduler/ExecutorServiceModule;", "", "()V", "createScheduler", "Ljava/util/concurrent/ExecutorService;", "app_debug"})
@dagger.hilt.InstallIn(value = {dagger.hilt.components.SingletonComponent.class})
public final class ExecutorServiceModule {
    @org.jetbrains.annotations.NotNull
    public static final org.csystem.android.app.generator.random.di.module.scheduler.ExecutorServiceModule INSTANCE = null;
    
    private ExecutorServiceModule() {
        super();
    }
    
    @dagger.Provides
    @javax.inject.Singleton
    @org.jetbrains.annotations.NotNull
    public final java.util.concurrent.ExecutorService createScheduler() {
        return null;
    }
}